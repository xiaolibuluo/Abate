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

    public interface IHandler24760
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24760(string Label, int Payload);

    public sealed class Handler24760 : IHandler24760
    {
        public void Handle(int value) { }
    }
    public static int Scale24761<T>(T value) where T : struct
    {
        int number = 93;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 28; }
        return number;
    }
    public static int Scale24762<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 39; }
        return number;
    }
    public sealed class SlotMap24763
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
    public static int Scale24764<T>(T value) where T : struct
    {
        int number = 17;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 61; }
        return number;
    }
    public static int Guard24765(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 24;
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
