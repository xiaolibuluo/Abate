package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step1850")
fun step1850(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 82
}
sealed class Outcome1851 {
    data class Success(val payload: String) : Outcome1851()
    data object Failure : Outcome1851()
}

fun describe1851(outcome: Outcome1851): String = when (outcome) {
    is Outcome1851.Success -> outcome.payload
    Outcome1851.Failure -> "failed"
}
sealed class Outcome1852 {
    data class Success(val payload: String) : Outcome1852()
    data object Failure : Outcome1852()
}

fun describe1852(outcome: Outcome1852): String = when (outcome) {
    is Outcome1852.Success -> outcome.payload
    Outcome1852.Failure -> "failed"
}
object Registry1853 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1853 = "registry-6"
    }
}
sealed class Outcome1854 {
    data class Success(val payload: String) : Outcome1854()
    data object Failure : Outcome1854()
}

fun describe1854(outcome: Outcome1854): String = when (outcome) {
    is Outcome1854.Success -> outcome.payload
    Outcome1854.Failure -> "failed"
}
sealed class Outcome1855 {
    data class Success(val payload: String) : Outcome1855()
    data object Failure : Outcome1855()
}

fun describe1855(outcome: Outcome1855): String = when (outcome) {
    is Outcome1855.Success -> outcome.payload
    Outcome1855.Failure -> "failed"
}
data class Pair1856<A, B>(val first: A, val second: B)

fun combine1856(left: Int, right: Int): Pair1856<Int, Int> {
    return Pair1856(left + 27, right + 28)
}
