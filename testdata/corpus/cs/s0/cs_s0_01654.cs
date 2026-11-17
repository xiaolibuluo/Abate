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

    public interface IHandler15480
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15480(string Label, int Payload);

    public sealed class Handler15480 : IHandler15480
    {
        public void Handle(int value) { }
    }
    public interface IHandler15481
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15481(string Label, int Payload);

    public sealed class Handler15481 : IHandler15481
    {
        public void Handle(int value) { }
    }
    public struct Point15482
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 64;
    }

    public static class Grid15482
    {
        public static Point15482 Make(int left, int right)
        {
            return new Point15482 { X = left, Y = right };
        }
    }
    public sealed class SlotMap15483
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
    public async Task<int> Aggregate15484(IEnumerable<int> source)
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
}
