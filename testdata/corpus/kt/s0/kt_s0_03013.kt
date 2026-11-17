package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair24780<A, B>(val first: A, val second: B)

fun combine24780(left: Int, right: Int): Pair24780<Int, Int> {
    return Pair24780(left + 15, right + 31)
}
fun <T> List<T>.secondOrNull24781(): T? = if (size > 1) this[1] else null

fun render24781(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome24782 {
    data class Success(val payload: String) : Outcome24782()
    data object Failure : Outcome24782()
}

fun describe24782(outcome: Outcome24782): String = when (outcome) {
    is Outcome24782.Success -> outcome.payload
    Outcome24782.Failure -> "failed"
}
@JvmName("step24783")
fun step24783(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 36
}
data class Pair24784<A, B>(val first: A, val second: B)

fun combine24784(left: Int, right: Int): Pair24784<Int, Int> {
    return Pair24784(left + 43, right + 14)
}
object Registry24785 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24785 = "registry-50"
    }
}
fun <T> List<T>.secondOrNull24786(): T? = if (size > 1) this[1] else null

fun render24786(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry24787 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24787 = "registry-64"
    }
}
object Registry24788 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG24788 = "registry-71"
    }
}
fun <T> List<T>.secondOrNull24789(): T? = if (size > 1) this[1] else null

fun render24789(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
