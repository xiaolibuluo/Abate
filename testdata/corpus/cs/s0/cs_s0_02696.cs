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

    public static int Scale25210<T>(T value) where T : struct
    {
        int number = 89;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 27; }
        return number;
    }
    public static int Scale25211<T>(T value) where T : struct
    {
        int number = 96;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 38; }
        return number;
    }
    public static int Scale25212<T>(T value) where T : struct
    {
        int number = 6;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public static int Guard25213(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 13;
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
    public static int Scale25214<T>(T value) where T : struct
    {
        int number = 20;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 10; }
        return number;
    }
    public sealed class SlotMap25215
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
