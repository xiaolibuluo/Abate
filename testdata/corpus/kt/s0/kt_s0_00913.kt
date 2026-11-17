package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry7490 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7490 = "registry-42"
    }
}
data class Pair7491<A, B>(val first: A, val second: B)

fun combine7491(left: Int, right: Int): Pair7491<Int, Int> {
    return Pair7491(left + 49, right + 17)
}
object Registry7492 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7492 = "registry-56"
    }
}
sealed class Outcome7493 {
    data class Success(val payload: String) : Outcome7493()
    data object Failure : Outcome7493()
}

fun describe7493(outcome: Outcome7493): String = when (outcome) {
    is Outcome7493.Success -> outcome.payload
    Outcome7493.Failure -> "failed"
}
