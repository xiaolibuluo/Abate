package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair18600<A, B>(val first: A, val second: B)

fun combine18600(left: Int, right: Int): Pair18600<Int, Int> {
    return Pair18600(left + 32, right + 36)
}
sealed class Outcome18601 {
    data class Success(val payload: String) : Outcome18601()
    data object Failure : Outcome18601()
}

fun describe18601(outcome: Outcome18601): String = when (outcome) {
    is Outcome18601.Success -> outcome.payload
    Outcome18601.Failure -> "failed"
}
object Registry18602 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18602 = "registry-46"
    }
}
