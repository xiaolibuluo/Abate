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

    public static int Scale20380<T>(T value) where T : struct
    {
        int number = 18;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 62; }
        return number;
    }
    public static int Guard20381(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 25;
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
    public sealed class SlotMap20382
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
    public struct Point20383
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 39;
    }

    public static class Grid20383
    {
        public static Point20383 Make(int left, int right)
        {
            return new Point20383 { X = left, Y = right };
        }
    }
    public interface IHandler20384
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20384(string Label, int Payload);

    public sealed class Handler20384 : IHandler20384
    {
        public void Handle(int value) { }
    }
}
