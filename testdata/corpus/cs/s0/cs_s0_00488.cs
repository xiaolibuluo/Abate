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

    public delegate void Notifier4580(int value);

    public sealed class Source4580
    {
        public event Notifier4580 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale4581<T>(T value) where T : struct
    {
        int number = 49;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 54; }
        return number;
    }
    public sealed class SlotMap4582
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
}
