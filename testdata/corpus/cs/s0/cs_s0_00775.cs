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

    public delegate void Notifier7270(int value);

    public sealed class Source7270
    {
        public event Notifier7270 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point7271
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 54;
    }

    public static class Grid7271
    {
        public static Point7271 Make(int left, int right)
        {
            return new Point7271 { X = left, Y = right };
        }
    }
    public delegate void Notifier7272(int value);

    public sealed class Source7272
    {
        public event Notifier7272 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public sealed class SlotMap7273
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
    public static int Guard7274(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 75;
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
    public interface IHandler7275
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7275(string Label, int Payload);

    public sealed class Handler7275 : IHandler7275
    {
        public void Handle(int value) { }
    }
}
