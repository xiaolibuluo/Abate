using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public delegate void Notifier20230(int value);

    public sealed class Source20230
    {
        public event Notifier20230 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier20231(int value);

    public sealed class Source20231
    {
        public event Notifier20231 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap20232
    {
        private readonly Dictionary<int, string> items = new Dictionary<int, string>();

        public string this[int index]
        {
            get => items.TryGetValue(index, out var value) ? value : string.Empty;
            set => items[index] = value;
        }

        public string Describe(int index) => index switch
        {
            0 => "zero",
            _ => items.ContainsKey(index) ? "present" : "absent"
        };
    }
    public delegate void Notifier20233(int value);

    public sealed class Source20233
    {
        public event Notifier20233 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
