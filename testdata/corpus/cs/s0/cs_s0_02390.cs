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

    public sealed class SlotMap22320
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
    public struct Point22321
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 25;
    }

    public static class Grid22321
    {
        public static Point22321 Make(int left, int right)
        {
            return new Point22321 { X = left, Y = right };
        }
    }
    public delegate void Notifier22322(int value);

    public sealed class Source22322
    {
        public event Notifier22322 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate22323(IEnumerable<int> source)
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
    public struct Point22324
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 46;
    }

    public static class Grid22324
    {
        public static Point22324 Make(int left, int right)
        {
            return new Point22324 { X = left, Y = right };
        }
    }
}
