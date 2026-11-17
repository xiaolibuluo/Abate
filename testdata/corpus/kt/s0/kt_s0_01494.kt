package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step12350")
fun step12350(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
sealed class Outcome12351 {
    data class Success(val payload: String) : Outcome12351()
    data object Failure : Outcome12351()
}

fun describe12351(outcome: Outcome12351): String = when (outcome) {
    is Outcome12351.Success -> outcome.payload
    Outcome12351.Failure -> "failed"
}
object Registry12352 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG12352 = "registry-69"
    }
}
data class Pair12353<A, B>(val first: A, val second: B)

fun combine12353(left: Int, right: Int): Pair12353<Int, Int> {
    return Pair12353(left + 76, right + 25)
}
