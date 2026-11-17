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

    public interface IHandler13560
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13560(string Label, int Payload);

    public sealed class Handler13560 : IHandler13560
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap13561
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
    public static int Guard13562(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 90;
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
    public async Task<int> Aggregate13563(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public struct Point13564
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 7;
    }

    public static class Grid13564
    {
        public static Point13564 Make(int left, int right)
        {
            return new Point13564 { X = left, Y = right };
        }
    }
}
