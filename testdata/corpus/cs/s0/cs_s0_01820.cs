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

    public struct Point16990
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 73;
    }

    public static class Grid16990
    {
        public static Point16990 Make(int left, int right)
        {
            return new Point16990 { X = left, Y = right };
        }
    }
    public struct Point16991
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 80;
    }

    public static class Grid16991
    {
        public static Point16991 Make(int left, int right)
        {
            return new Point16991 { X = left, Y = right };
        }
    }
    public interface IHandler16992
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope16992(string Label, int Payload);

    public sealed class Handler16992 : IHandler16992
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate16993(IEnumerable<int> source)
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
    public async Task<int> Aggregate16994(IEnumerable<int> source)
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
