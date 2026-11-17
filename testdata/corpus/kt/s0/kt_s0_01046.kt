package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry8600 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8600 = "registry-30"
    }
}
sealed class Outcome8601 {
    data class Success(val payload: String) : Outcome8601()
    data object Failure : Outcome8601()
}

fun describe8601(outcome: Outcome8601): String = when (outcome) {
    is Outcome8601.Success -> outcome.payload
    Outcome8601.Failure -> "failed"
}
@JvmName("step8602")
fun step8602(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 44
}
sealed class Outcome8603 {
    data class Success(val payload: String) : Outcome8603()
    data object Failure : Outcome8603()
}

fun describe8603(outcome: Outcome8603): String = when (outcome) {
    is Outcome8603.Success -> outcome.payload
    Outcome8603.Failure -> "failed"
}
data class Pair8604<A, B>(val first: A, val second: B)

fun combine8604(left: Int, right: Int): Pair8604<Int, Int> {
    return Pair8604(left + 58, right + 34)
}
data class Pair8605<A, B>(val first: A, val second: B)

fun combine8605(left: Int, right: Int): Pair8605<Int, Int> {
    return Pair8605(left + 65, right + 45)
}
@JvmName("step8606")
fun step8606(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 72
}
