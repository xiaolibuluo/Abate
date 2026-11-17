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

    public static int Guard15990(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 34;
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
    public interface IHandler15991
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15991(string Label, int Payload);

    public sealed class Handler15991 : IHandler15991
    {
        public void Handle(int value) { }
    }
    public struct Point15992
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 48;
    }

    public static class Grid15992
    {
        public static Point15992 Make(int left, int right)
        {
            return new Point15992 { X = left, Y = right };
        }
    }
    public static int Guard15993(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 55;
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
    public sealed class SlotMap15994
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
