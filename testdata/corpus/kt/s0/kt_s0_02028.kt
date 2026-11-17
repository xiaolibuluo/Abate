package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16770")
fun step16770(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 78
}
data class Pair16771<A, B>(val first: A, val second: B)

fun combine16771(left: Int, right: Int): Pair16771<Int, Int> {
    return Pair16771(left + 85, right + 47)
}
sealed class Outcome16772 {
    data class Success(val payload: String) : Outcome16772()
    data object Failure : Outcome16772()
}

fun describe16772(outcome: Outcome16772): String = when (outcome) {
    is Outcome16772.Success -> outcome.payload
    Outcome16772.Failure -> "failed"
}
data class Pair16773<A, B>(val first: A, val second: B)

fun combine16773(left: Int, right: Int): Pair16773<Int, Int> {
    return Pair16773(left + 99, right + 69)
}
fun <T> List<T>.secondOrNull16774(): T? = if (size > 1) this[1] else null

fun render16774(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
