package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step16480")
fun step16480(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
sealed class Outcome16481 {
    data class Success(val payload: String) : Outcome16481()
    data object Failure : Outcome16481()
}

fun describe16481(outcome: Outcome16481): String = when (outcome) {
    is Outcome16481.Success -> outcome.payload
    Outcome16481.Failure -> "failed"
}
object Registry16482 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16482 = "registry-6"
    }
}
fun <T> List<T>.secondOrNull16483(): T? = if (size > 1) this[1] else null

fun render16483(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16484 {
    data class Success(val payload: String) : Outcome16484()
    data object Failure : Outcome16484()
}

fun describe16484(outcome: Outcome16484): String = when (outcome) {
    is Outcome16484.Success -> outcome.payload
    Outcome16484.Failure -> "failed"
}
fun <T> List<T>.secondOrNull16485(): T? = if (size > 1) this[1] else null

fun render16485(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step16486")
fun step16486(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 34
}
fun <T> List<T>.secondOrNull16487(): T? = if (size > 1) this[1] else null

fun render16487(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
