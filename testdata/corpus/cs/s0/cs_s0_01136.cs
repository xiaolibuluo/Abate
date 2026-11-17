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

    public struct Point10650
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 76;
    }

    public static class Grid10650
    {
        public static Point10650 Make(int left, int right)
        {
            return new Point10650 { X = left, Y = right };
        }
    }
    public interface IHandler10651
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10651(string Label, int Payload);

    public sealed class Handler10651 : IHandler10651
    {
        public void Handle(int value) { }
    }
    public interface IHandler10652
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope10652(string Label, int Payload);

    public sealed class Handler10652 : IHandler10652
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate10653(IEnumerable<int> source)
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
    public sealed class SlotMap10654
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
