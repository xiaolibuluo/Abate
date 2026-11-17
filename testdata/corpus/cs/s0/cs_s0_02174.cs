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

    public static int Guard20290(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 95;
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
    public struct Point20291
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 5;
    }

    public static class Grid20291
    {
        public static Point20291 Make(int left, int right)
        {
            return new Point20291 { X = left, Y = right };
        }
    }
    public delegate void Notifier20292(int value);

    public sealed class Source20292
    {
        public event Notifier20292 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler20293
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope20293(string Label, int Payload);

    public sealed class Handler20293 : IHandler20293
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap20294
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
