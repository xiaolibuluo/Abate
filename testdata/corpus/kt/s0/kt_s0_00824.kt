package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull6790(): T? = if (size > 1) this[1] else null

fun render6790(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome6791 {
    data class Success(val payload: String) : Outcome6791()
    data object Failure : Outcome6791()
}

fun describe6791(outcome: Outcome6791): String = when (outcome) {
    is Outcome6791.Success -> outcome.payload
    Outcome6791.Failure -> "failed"
}
data class Pair6792<A, B>(val first: A, val second: B)

fun combine6792(left: Int, right: Int): Pair6792<Int, Int> {
    return Pair6792(left + 19, right + 26)
}
object Registry6793 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6793 = "registry-26"
    }
}
fun <T> List<T>.secondOrNull6794(): T? = if (size > 1) this[1] else null

fun render6794(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
sealed class Outcome6795 {
    data class Success(val payload: String) : Outcome6795()
    data object Failure : Outcome6795()
}

fun describe6795(outcome: Outcome6795): String = when (outcome) {
    is Outcome6795.Success -> outcome.payload
    Outcome6795.Failure -> "failed"
}
object Registry6796 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6796 = "registry-47"
    }
}
@JvmName("step6797")
fun step6797(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 54
}
fun <T> List<T>.secondOrNull6798(): T? = if (size > 1) this[1] else null

fun render6798(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
