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

    public struct Point17370
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 82;
    }

    public static class Grid17370
    {
        public static Point17370 Make(int left, int right)
        {
            return new Point17370 { X = left, Y = right };
        }
    }
    public static int Guard17371(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 89;
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
    public sealed class SlotMap17372
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
    public interface IHandler17373
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17373(string Label, int Payload);

    public sealed class Handler17373 : IHandler17373
    {
        public void Handle(int value) { }
    }
}
