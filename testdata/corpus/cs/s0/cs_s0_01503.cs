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

    public struct Point14070
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid14070
    {
        public static Point14070 Make(int left, int right)
        {
            return new Point14070 { X = left, Y = right };
        }
    }
    public interface IHandler14071
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14071(string Label, int Payload);

    public sealed class Handler14071 : IHandler14071
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier14072(int value);

    public sealed class Source14072
    {
        public event Notifier14072 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard14073(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 81;
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
    public sealed class SlotMap14074
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
