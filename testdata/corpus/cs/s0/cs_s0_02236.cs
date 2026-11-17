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

    public interface IHandler20870
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20870(string Label, int Payload);

    public sealed class Handler20870 : IHandler20870
    {
        public void Handle(int value) { }
    }
    public static int Scale20871<T>(T value) where T : struct
    {
        int number = 80;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 50; }
        return number;
    }
    public sealed class SlotMap20872
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
    public static int Guard20873(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 94;
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
    public struct Point20874
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 101;
    }

    public static class Grid20874
    {
        public static Point20874 Make(int left, int right)
        {
            return new Point20874 { X = left, Y = right };
        }
    }
}
