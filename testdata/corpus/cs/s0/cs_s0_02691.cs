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

    public struct Point25160
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 24;
    }

    public static class Grid25160
    {
        public static Point25160 Make(int left, int right)
        {
            return new Point25160 { X = left, Y = right };
        }
    }
    public sealed class SlotMap25161
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
    public static int Guard25162(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 38;
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
    public static int Scale25163<T>(T value) where T : struct
    {
        int number = 45;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
}
