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

    public interface IHandler25230
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25230(string Label, int Payload);

    public sealed class Handler25230 : IHandler25230
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier25231(int value);

    public sealed class Source25231
    {
        public event Notifier25231 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point25232
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 32;
    }

    public static class Grid25232
    {
        public static Point25232 Make(int left, int right)
        {
            return new Point25232 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate25233(IEnumerable<int> source)
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
    public static int Guard25234(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 46;
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
