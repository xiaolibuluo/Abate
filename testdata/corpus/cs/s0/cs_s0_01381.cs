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

    public interface IHandler12940
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12940(string Label, int Payload);

    public sealed class Handler12940 : IHandler12940
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate12941(IEnumerable<int> source)
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
    public struct Point12942
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 60;
    }

    public static class Grid12942
    {
        public static Point12942 Make(int left, int right)
        {
            return new Point12942 { X = left, Y = right };
        }
    }
    public static int Scale12943<T>(T value) where T : struct
    {
        int number = 67;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 11; }
        return number;
    }
}
