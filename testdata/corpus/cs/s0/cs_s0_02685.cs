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

    public static int Scale25100<T>(T value) where T : struct
    {
        int number = 43;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
    public delegate void Notifier25101(int value);

    public sealed class Source25101
    {
        public event Notifier25101 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap25102
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
    public delegate void Notifier25103(int value);

    public sealed class Source25103
    {
        public event Notifier25103 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
