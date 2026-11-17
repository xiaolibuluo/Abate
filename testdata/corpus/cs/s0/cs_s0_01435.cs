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

    public delegate void Notifier13440(int value);

    public sealed class Source13440
    {
        public event Notifier13440 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap13441
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
    public delegate void Notifier13442(int value);

    public sealed class Source13442
    {
        public event Notifier13442 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier13443(int value);

    public sealed class Source13443
    {
        public event Notifier13443 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
