export default {
  language: 'Мова',
  ddos: {
    description: 'Атака DDoS робить величезну кількість запитів на ворожі сайти, чим спричиняє їхню величезну завантаженість і виводить їх з ладу. В наслідку атаки DDoS, сервери не можуть працювати в штатному режимі / лагають, або взагалі не працюють. Деталі тут: https://help-ukraine-win.com/',
    coordinators: 'Для координацій цілей: https://discord.gg/7BfJ9JKQ98 ',
    counter: {
      attackedTimes: 'Атаковано разів:',
      currentTarget: 'Теперішня ціль: '
    },
    enable: {
      name: 'Атака DDoS',
      description: 'Вмикає атаку DDoS по основним системам ворога: урядові сайти, банківські системи, пропагандиські пабліки, тощо. Поточні цілі атаки відображаються на сторінці'
    },
    proxy: {
      name: 'Атакувати тільки через проксі',
      description: 'З атакою без проксі, програма зможе використовувати ваш реальний IP. Це дозволяє робити атаки більш ефективними, але водночас зробить ваш IP відомим ворогу'
    },
    advanced: {
      header: 'Розширені налаштування',
      description: 'Якщо ви недосвідчений користувач і не розумієте частини понять, будь ласка не змінюйте розширених налаштувань. Це може навпаки погано вплинути на роботу програми',
      masDosersCount: {
        name: 'Рівень паралельності',
        description: 'Задає кількість паралельних процесів. Використовуються асинхронні механізми (не потокові), тому це не сильно впливає на CPU/RAM. Після зміни, треба почекати ~10 cекунд поки стартують додаткові паралельні досери'
      }
    },
    update: {
      title: 'Оновлення доступне',
      cancel: 'Відміна',
      confirm: 'Встановити оновлення'
    }
  }
}
