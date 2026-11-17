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

    public interface IHandler26240
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope26240(string Label, int Payload);

    public sealed class Handler26240 : IHandler26240
    {
        public void Handle(int value) { }
    }
    public static int Guard26241(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 77;
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
    public struct Point26242
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 84;
    }

    public static class Grid26242
    {
        public static Point26242 Make(int left, int right)
        {
            return new Point26242 { X = left, Y = right };
        }
    }
    public sealed class SlotMap26243
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
    public delegate void Notifier26244(int value);

    public sealed class Source26244
    {
        public event Notifier26244 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
