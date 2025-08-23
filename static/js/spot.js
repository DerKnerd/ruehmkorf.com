import { get, put } from '../lib/jinya-http.js';
import alert from '../lib/ui/alert.js';
import { parse } from '../lib/csv-parser.js';

Alpine.data('spotData', () => ({
  loading: true,
  mappings: [],
  async init() {
    this.loading = true;
    this.mappings = await get('/api/spot/mapping');
    this.loading = false;
  },
  loadSpot() {
    const filePicker = document.createElement('input');
    filePicker.type = 'file';
    filePicker.accept = '.csv';
    filePicker.multiple = false;
    filePicker.addEventListener('change', async (e) => {
      const file = e.target.files[0];
      if (file) {
        const reader = new FileReader();
        reader.onload = async (e) => {
          const csv = e.target.result;
          const parsed = parse(csv);
          this.mappings = parsed
            .slice(1)
            .filter((f) => f.length === 3)
            .map((line) => {
              let character = line[0];
              if (character.length > 1 && character.startsWith("'")) {
                character = character.substring(1, 2);
              }

              return {
                character,
                english: line[1],
                german: line[2],
              };
            });
        };
        reader.readAsText(file);
      }
      filePicker.remove();
    });
    filePicker.style.display = 'none';
    document.body.appendChild(filePicker);
    filePicker.click();
  },
  async saveSpot() {
    try {
      await put('/api/spot/mapping', this.mappings);
      alert({
        title: 'Spell-O-Tron updated',
        message: 'The Spell-O-Tron has been updated successfully.',
        closeLabel: 'Alright',
        positive: true,
      });
    } catch (e) {
      alert({
        title: 'Error updating Spell-O-Tron',
        message: 'The Spell-O-Tron could not be updated. Please contact support.',
        closeLabel: 'Okay',
        negative: true,
      });
    }
  },
}));
