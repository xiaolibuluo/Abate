package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18660<A, B>(val first: A, val second: B)

fun combine18660(left: Int, right: Int): Pair18660<Int, Int> {
    return Pair18660(left + 13, right + 17)
}
object Registry18661 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18661 = "registry-20"
    }
}
sealed class Outcome18662 {
    data class Success(val payload: String) : Outcome18662()
    data object Failure : Outcome18662()
}

fun describe18662(outcome: Outcome18662): String = when (outcome) {
    is Outcome18662.Success -> outcome.payload
    Outcome18662.Failure -> "failed"
}
data class Pair18663<A, B>(val first: A, val second: B)

fun combine18663(left: Int, right: Int): Pair18663<Int, Int> {
    return Pair18663(left + 34, right + 50)
}
