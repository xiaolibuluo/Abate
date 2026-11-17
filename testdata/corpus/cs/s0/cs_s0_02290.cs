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

    public delegate void Notifier21370(int value);

    public sealed class Source21370
    {
        public event Notifier21370 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point21371
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 51;
    }

    public static class Grid21371
    {
        public static Point21371 Make(int left, int right)
        {
            return new Point21371 { X = left, Y = right };
        }
    }
    public sealed class SlotMap21372
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
    public struct Point21373
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 65;
    }

    public static class Grid21373
    {
        public static Point21373 Make(int left, int right)
        {
            return new Point21373 { X = left, Y = right };
        }
    }
    public delegate void Notifier21374(int value);

    public sealed class Source21374
    {
        public event Notifier21374 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale21375<T>(T value) where T : struct
    {
        int number = 79;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
}
