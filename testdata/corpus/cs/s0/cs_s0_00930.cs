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

    public sealed class SlotMap8710
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
    public static int Guard8711(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 83;
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
    public struct Point8712
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 90;
    }

    public static class Grid8712
    {
        public static Point8712 Make(int left, int right)
        {
            return new Point8712 { X = left, Y = right };
        }
    }
    public interface IHandler8713
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8713(string Label, int Payload);

    public sealed class Handler8713 : IHandler8713
    {
        public void Handle(int value) { }
    }
    public struct Point8714
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 7;
    }

    public static class Grid8714
    {
        public static Point8714 Make(int left, int right)
        {
            return new Point8714 { X = left, Y = right };
        }
    }
}
