package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step6710")
fun step6710(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 95
}
@JvmName("step6711")
fun step6711(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 5
}
sealed class Outcome6712 {
    data class Success(val payload: String) : Outcome6712()
    data object Failure : Outcome6712()
}

fun describe6712(outcome: Outcome6712): String = when (outcome) {
    is Outcome6712.Success -> outcome.payload
    Outcome6712.Failure -> "failed"
}
sealed class Outcome6713 {
    data class Success(val payload: String) : Outcome6713()
    data object Failure : Outcome6713()
}

fun describe6713(outcome: Outcome6713): String = when (outcome) {
    is Outcome6713.Success -> outcome.payload
    Outcome6713.Failure -> "failed"
}
@JvmName("step6714")
fun step6714(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 26
}
@JvmName("step6715")
fun step6715(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 33
}
data class Pair6716<A, B>(val first: A, val second: B)

fun combine6716(left: Int, right: Int): Pair6716<Int, Int> {
    return Pair6716(left + 40, right + 14)
}
fun <T> List<T>.secondOrNull6717(): T? = if (size > 1) this[1] else null

fun render6717(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
