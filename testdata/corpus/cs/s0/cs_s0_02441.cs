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

    public interface IHandler22800
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22800(string Label, int Payload);

    public sealed class Handler22800 : IHandler22800
    {
        public void Handle(int value) { }
    }
    public interface IHandler22801
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22801(string Label, int Payload);

    public sealed class Handler22801 : IHandler22801
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap22802
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
    public struct Point22803
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 81;
    }

    public static class Grid22803
    {
        public static Point22803 Make(int left, int right)
        {
            return new Point22803 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate22804(IEnumerable<int> source)
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
