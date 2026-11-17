package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

sealed class Outcome21560 {
    data class Success(val payload: String) : Outcome21560()
    data object Failure : Outcome21560()
}

fun describe21560(outcome: Outcome21560): String = when (outcome) {
    is Outcome21560.Success -> outcome.payload
    Outcome21560.Failure -> "failed"
}
data class Pair21561<A, B>(val first: A, val second: B)

fun combine21561(left: Int, right: Int): Pair21561<Int, Int> {
    return Pair21561(left + 7, right + 45)
}
object Registry21562 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21562 = "registry-14"
    }
}
data class Pair21563<A, B>(val first: A, val second: B)

fun combine21563(left: Int, right: Int): Pair21563<Int, Int> {
    return Pair21563(left + 21, right + 67)
}
