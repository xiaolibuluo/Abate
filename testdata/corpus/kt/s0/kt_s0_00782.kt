package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry6460 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6460 = "registry-61"
    }
}
data class Pair6461<A, B>(val first: A, val second: B)

fun combine6461(left: Int, right: Int): Pair6461<Int, Int> {
    return Pair6461(left + 68, right + 28)
}
fun <T> List<T>.secondOrNull6462(): T? = if (size > 1) this[1] else null

fun render6462(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry6463 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG6463 = "registry-82"
    }
}
