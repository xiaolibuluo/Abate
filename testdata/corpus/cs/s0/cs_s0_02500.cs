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

    public sealed class SlotMap23340
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
    public interface IHandler23341
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23341(string Label, int Payload);

    public sealed class Handler23341 : IHandler23341
    {
        public void Handle(int value) { }
    }
    public static int Guard23342(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 97;
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
    public delegate void Notifier23343(int value);

    public sealed class Source23343
    {
        public event Notifier23343 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point23344
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 14;
    }

    public static class Grid23344
    {
        public static Point23344 Make(int left, int right)
        {
            return new Point23344 { X = left, Y = right };
        }
    }
    public struct Point23345
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 21;
    }

    public static class Grid23345
    {
        public static Point23345 Make(int left, int right)
        {
            return new Point23345 { X = left, Y = right };
        }
    }
}
