package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair1760<A, B>(val first: A, val second: B)

fun combine1760(left: Int, right: Int): Pair1760<Int, Int> {
    return Pair1760(left + 62, right + 21)
}
object Registry1761 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1761 = "registry-69"
    }
}
fun <T> List<T>.secondOrNull1762(): T? = if (size > 1) this[1] else null

fun render1762(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry1763 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1763 = "registry-83"
    }
}
data class Pair1764<A, B>(val first: A, val second: B)

fun combine1764(left: Int, right: Int): Pair1764<Int, Int> {
    return Pair1764(left + 90, right + 65)
}
sealed class Outcome1765 {
    data class Success(val payload: String) : Outcome1765()
    data object Failure : Outcome1765()
}

fun describe1765(outcome: Outcome1765): String = when (outcome) {
    is Outcome1765.Success -> outcome.payload
    Outcome1765.Failure -> "failed"
}
@JvmName("step1766")
fun step1766(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 7
}
