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

    public struct Point20770
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 40;
    }

    public static class Grid20770
    {
        public static Point20770 Make(int left, int right)
        {
            return new Point20770 { X = left, Y = right };
        }
    }
    public sealed class SlotMap20771
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
    public static int Scale20772<T>(T value) where T : struct
    {
        int number = 54;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public delegate void Notifier20773(int value);

    public sealed class Source20773
    {
        public event Notifier20773 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
