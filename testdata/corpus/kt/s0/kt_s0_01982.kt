package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16380<A, B>(val first: A, val second: B)

fun combine16380(left: Int, right: Int): Pair16380<Int, Int> {
    return Pair16380(left + 56, right + 68)
}
fun <T> List<T>.secondOrNull16381(): T? = if (size > 1) this[1] else null

fun render16381(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull16382(): T? = if (size > 1) this[1] else null

fun render16382(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome16383 {
    data class Success(val payload: String) : Outcome16383()
    data object Failure : Outcome16383()
}

fun describe16383(outcome: Outcome16383): String = when (outcome) {
    is Outcome16383.Success -> outcome.payload
    Outcome16383.Failure -> "failed"
}
object Registry16384 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16384 = "registry-84"
    }
}
data class Pair16385<A, B>(val first: A, val second: B)

fun combine16385(left: Int, right: Int): Pair16385<Int, Int> {
    return Pair16385(left + 91, right + 62)
}
@JvmName("step16386")
fun step16386(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 98
}
