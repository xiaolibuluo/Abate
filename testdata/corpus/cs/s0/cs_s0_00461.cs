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

    public async Task<int> Aggregate4320(IEnumerable<int> source)
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
    public sealed class SlotMap4321
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
    public struct Point4322
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 9;
    }

    public static class Grid4322
    {
        public static Point4322 Make(int left, int right)
        {
            return new Point4322 { X = left, Y = right };
        }
    }
    public delegate void Notifier4323(int value);

    public sealed class Source4323
    {
        public event Notifier4323 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler4324
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4324(string Label, int Payload);

    public sealed class Handler4324 : IHandler4324
    {
        public void Handle(int value) { }
    }
}
