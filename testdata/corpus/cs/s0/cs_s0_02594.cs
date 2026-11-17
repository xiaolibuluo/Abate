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

    public sealed class SlotMap24260
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
    public static int Scale24261<T>(T value) where T : struct
    {
        int number = 25;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public struct Point24262
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 32;
    }

    public static class Grid24262
    {
        public static Point24262 Make(int left, int right)
        {
            return new Point24262 { X = left, Y = right };
        }
    }
    public delegate void Notifier24263(int value);

    public sealed class Source24263
    {
        public event Notifier24263 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard24264(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 46;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
}
