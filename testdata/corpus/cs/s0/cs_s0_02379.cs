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

    public delegate void Notifier22210(int value);

    public sealed class Source22210
    {
        public event Notifier22210 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point22211
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 76;
    }

    public static class Grid22211
    {
        public static Point22211 Make(int left, int right)
        {
            return new Point22211 { X = left, Y = right };
        }
    }
    public delegate void Notifier22212(int value);

    public sealed class Source22212
    {
        public event Notifier22212 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale22213<T>(T value) where T : struct
    {
        int number = 90;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public sealed class SlotMap22214
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
