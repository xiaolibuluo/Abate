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

    public delegate void Notifier22300(int value);

    public sealed class Source22300
    {
        public event Notifier22300 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point22301
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 96;
    }

    public static class Grid22301
    {
        public static Point22301 Make(int left, int right)
        {
            return new Point22301 { X = left, Y = right };
        }
    }
    public interface IHandler22302
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22302(string Label, int Payload);

    public sealed class Handler22302 : IHandler22302
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate22303(IEnumerable<int> source)
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
    public sealed class SlotMap22304
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
