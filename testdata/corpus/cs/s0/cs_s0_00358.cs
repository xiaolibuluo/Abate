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

    public async Task<int> Aggregate3390(IEnumerable<int> source)
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
    public struct Point3391
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 54;
    }

    public static class Grid3391
    {
        public static Point3391 Make(int left, int right)
        {
            return new Point3391 { X = left, Y = right };
        }
    }
    public sealed class SlotMap3392
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
    public interface IHandler3393
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3393(string Label, int Payload);

    public sealed class Handler3393 : IHandler3393
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier3394(int value);

    public sealed class Source3394
    {
        public event Notifier3394 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
