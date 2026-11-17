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

    public sealed class SlotMap18940
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
    public static int Scale18941<T>(T value) where T : struct
    {
        int number = 93;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 41; }
        return number;
    }
    public static int Scale18942<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public struct Point18943
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 10;
    }

    public static class Grid18943
    {
        public static Point18943 Make(int left, int right)
        {
            return new Point18943 { X = left, Y = right };
        }
    }
    public delegate void Notifier18944(int value);

    public sealed class Source18944
    {
        public event Notifier18944 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale18945<T>(T value) where T : struct
    {
        int number = 24;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 24; }
        return number;
    }
}
