package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24030<A, B>(val first: A, val second: B)

fun combine24030(left: Int, right: Int): Pair24030<Int, Int> {
    return Pair24030(left + 10, right + 55)
}
fun <T> List<T>.secondOrNull24031(): T? = if (size > 1) this[1] else null

fun render24031(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step24032")
fun step24032(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 24
}
data class Pair24033<A, B>(val first: A, val second: B)

fun combine24033(left: Int, right: Int): Pair24033<Int, Int> {
    return Pair24033(left + 31, right + 27)
}
sealed class Outcome24034 {
    data class Success(val payload: String) : Outcome24034()
    data object Failure : Outcome24034()
}

fun describe24034(outcome: Outcome24034): String = when (outcome) {
    is Outcome24034.Success -> outcome.payload
    Outcome24034.Failure -> "failed"
}
