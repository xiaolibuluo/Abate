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

    public delegate void Notifier25140(int value);

    public sealed class Source25140
    {
        public event Notifier25140 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale25141<T>(T value) where T : struct
    {
        int number = 5;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public delegate void Notifier25142(int value);

    public sealed class Source25142
    {
        public event Notifier25142 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale25143<T>(T value) where T : struct
    {
        int number = 19;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
    public sealed class SlotMap25144
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
