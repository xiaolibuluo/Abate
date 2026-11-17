package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step7750")
fun step7750(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 89
}
data class Pair7751<A, B>(val first: A, val second: B)

fun combine7751(left: Int, right: Int): Pair7751<Int, Int> {
    return Pair7751(left + 96, right + 16)
}
fun <T> List<T>.secondOrNull7752(): T? = if (size > 1) this[1] else null

fun render7752(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair7753<A, B>(val first: A, val second: B)

fun combine7753(left: Int, right: Int): Pair7753<Int, Int> {
    return Pair7753(left + 13, right + 38)
}
fun <T, R> List<T>.mapChecked7754(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step7755")
fun step7755(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 27
}
fun <T, R> List<T>.mapChecked7756(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
sealed class Outcome7757 {
    data class Success(val payload: String) : Outcome7757()
    data object Failure : Outcome7757()
}

fun describe7757(outcome: Outcome7757): String = when (outcome) {
    is Outcome7757.Success -> outcome.payload
    Outcome7757.Failure -> "failed"
}
@JvmName("step7758")
fun step7758(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 48
}
object Registry7759 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG7759 = "registry-55"
    }
}
