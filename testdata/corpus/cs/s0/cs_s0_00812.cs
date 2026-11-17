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

    public static int Guard7620(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 17;
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
    public sealed class SlotMap7621
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
    public interface IHandler7622
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope7622(string Label, int Payload);

    public sealed class Handler7622 : IHandler7622
    {
        public void Handle(int value) { }
    }
    public struct Point7623
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 38;
    }

    public static class Grid7623
    {
        public static Point7623 Make(int left, int right)
        {
            return new Point7623 { X = left, Y = right };
        }
    }
}
