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

    public static int Guard23740(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 21;
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
    public interface IHandler23741
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23741(string Label, int Payload);

    public sealed class Handler23741 : IHandler23741
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap23742
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
    public struct Point23743
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 42;
    }

    public static class Grid23743
    {
        public static Point23743 Make(int left, int right)
        {
            return new Point23743 { X = left, Y = right };
        }
    }
}
