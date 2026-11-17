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

    public struct Point6750
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 50;
    }

    public static class Grid6750
    {
        public static Point6750 Make(int left, int right)
        {
            return new Point6750 { X = left, Y = right };
        }
    }
    public static int Guard6751(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 57;
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
    public static int Guard6752(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 64;
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
    public interface IHandler6753
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope6753(string Label, int Payload);

    public sealed class Handler6753 : IHandler6753
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap6754
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
