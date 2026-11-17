package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair7080<A, B>(val first: A, val second: B)

fun combine7080(left: Int, right: Int): Pair7080<Int, Int> {
    return Pair7080(left + 91, right + 24)
}
data class Pair7081<A, B>(val first: A, val second: B)

fun combine7081(left: Int, right: Int): Pair7081<Int, Int> {
    return Pair7081(left + 98, right + 35)
}
object Registry7082 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7082 = "registry-8"
    }
}
@JvmName("step7083")
fun step7083(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 15
}
fun <T, R> List<T>.mapChecked7084(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome7085 {
    data class Success(val payload: String) : Outcome7085()
    data object Failure : Outcome7085()
}

fun describe7085(outcome: Outcome7085): String = when (outcome) {
    is Outcome7085.Success -> outcome.payload
    Outcome7085.Failure -> "failed"
}
