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

    public static int Scale23190<T>(T value) where T : struct
    {
        int number = 82;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
    public delegate void Notifier23191(int value);

    public sealed class Source23191
    {
        public event Notifier23191 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier23192(int value);

    public sealed class Source23192
    {
        public event Notifier23192 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap23193
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
    public static int Scale23194<T>(T value) where T : struct
    {
        int number = 13;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 60; }
        return number;
    }
}
