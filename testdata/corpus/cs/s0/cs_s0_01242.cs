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

    public interface IHandler11640
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11640(string Label, int Payload);

    public sealed class Handler11640 : IHandler11640
    {
        public void Handle(int value) { }
    }
    public struct Point11641
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 12;
    }

    public static class Grid11641
    {
        public static Point11641 Make(int left, int right)
        {
            return new Point11641 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate11642(IEnumerable<int> source)
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
    public sealed class SlotMap11643
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
    public static int Guard11644(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 33;
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
}
