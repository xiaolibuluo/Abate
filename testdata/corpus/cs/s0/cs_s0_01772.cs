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

    public static int Guard16520(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 44;
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
    public interface IHandler16521
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16521(string Label, int Payload);

    public sealed class Handler16521 : IHandler16521
    {
        public void Handle(int value) { }
    }
    public struct Point16522
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 58;
    }

    public static class Grid16522
    {
        public static Point16522 Make(int left, int right)
        {
            return new Point16522 { X = left, Y = right };
        }
    }
    public static int Scale16523<T>(T value) where T : struct
    {
        int number = 65;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
    public sealed class SlotMap16524
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
